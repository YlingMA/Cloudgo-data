
import (
	"testing"
)

func TestUser(t *testing.T) {
	// Clear the table.
	UserService.DeleteAll()

	// Create new users.
	testUser1 := NewUserInfo("myl", "zhongda")
	testUser2 := NewUserInfo("mayiling", "ZhongDa")

	// Insert the users.
	UserService.Insert(testUser1)
	UserService.Insert(testUser2)

	// Test FindAll
	allUser := UserService.FindAll()
	if len(allUser) != 2 {
		t.Error("Expected length 2, got ", len(allUser))
	}

	// Test FindByID
	findUser1 := UserService.FindByID(testUser1.UID)
	if findUser1.Username != "Alice" {
		t.Error("Expected ", testUser1.Username, ", got ", findUser1.Username)
	}
	findUser2 := UserService.FindByID(testUser2.UID)
	if findUser2.Username != "Bob" {
		t.Error("Expected ", testUser2.Username, ", got ", findUser2.Username)
	}
}
