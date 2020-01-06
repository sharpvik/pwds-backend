package antidos

import (
	"time"
	"net/http"

	"github.com/sharpvik/pwds-backend/iputils"
)



// AntiDoS struct allows one to track IPs that send server requests and limit
// number of those requests so as to prevent basic DoS attacks.
type AntiDoS struct {
	maxAttempts		int
	banDuration 	time.Duration

	noticedIPs 		map[string]int
	blockedIPs 		map[string]time.Time
}



// New returns pointer to a newly initialized instance of AntiDoS based on
// maxAttempts and banDuration.
func New(maxAttempts int, banDuration time.Duration) *AntiDoS {
	return &AntiDoS{
		maxAttempts,
		banDuration,
		make(map[string]int),
		make(map[string]time.Time),
	}
}



// Notice method brings new IP addresses to AntiDoS's attention. It reutrns true
// or false to signify whether this IP should be served or is banned.
func (ads *AntiDoS) Notice(r *http.Request) bool {
	ip := iputils.ReadCleanIP(r)

	numberOfAttemptsLeft, noticed := ads.noticedIPs[ip]
	momentUntilBlocked, blocked := ads.blockedIPs[ip]

	// If this ip is already blocked ...
	if blocked {
		// ... check whether it's time to remove ban.
		unban := time.Now().Sub(momentUntilBlocked).Nanoseconds() >= 0

		// If not, they are still banned!
		if !unban {
			return false
		}

		// If yes, remove ip from blockedIPs list
		delete(ads.blockedIPs, ip)
	}

	// If this ip has not been noticed yet and is not blocked (anymore) ...
	if !noticed {
		// ... set their number of attempts to (maxAttempts - 1) and store their
		// ip in the list.
		ads.noticedIPs[ip] = ads.maxAttempts - 1
		return true
	}

	// While they haven't used up all their attempts, decrement counter.
	if numberOfAttemptsLeft > 0 {
		ads.noticedIPs[ip]--
		return true
	}

	// If they are noticed and they ran out of attempts for now, block them.
	delete(ads.noticedIPs, ip)
	ads.blockedIPs[ip] = time.Now().Add(ads.banDuration)
	return false
}
