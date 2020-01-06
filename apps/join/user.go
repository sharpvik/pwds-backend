package join

import (
	"os"
	"path"
	"fmt"
	
	"github.com/sharpvik/pwds-backend/config"
)



// User type reflects JSON input from the client side.
// JSON data is parsed into its instance and stored.
type User struct {
	Username	string
	Masterhash	string
}



// NewUser function returns pointer to an empty User instance.
// Its attributes are not initialized in order to fill them with JSON data.
func NewUser() *User {
	return &User{}
}



// Store function uses paths specified in the config file to store user data.
func (u *User) Store() error {
	folder := path.Join(config.UserStoreFolder, u.Username)
	err := os.Mkdir(folder, 0777)

	if err != nil {
		fmt.Println(err)
		return err
	}

	masterhashPath := path.Join(folder, ".masterhash")
	f, err := os.Create(masterhashPath)

	if err != nil {
		fmt.Println(err)
		return err
	}

	_, err = f.Write( []byte(u.Masterhash) )

	if err != nil {
		fmt.Println(err)
	}

	return err
}
