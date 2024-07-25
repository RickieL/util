package sql

import (
	"reflect"
	"testing"
)

func TestSplitSQL(t *testing.T) {
	sql1 := "SELECT * FROM users; INSERT INTO orders (user_id, product_id) VALUES (1, 2)"
	expected1 := []string{"SELECT * FROM users;", "INSERT INTO orders (user_id, product_id) VALUES (1, 2)"}
	result := SplitSQL(sql1)
	if !reflect.DeepEqual(result, expected1) {
		t.Errorf("SplitSQL() = %v, expected %v", result, expected1)
	}

	sql2 := `SELECT * FROM users; 
	INSERT INTO orders (user_id, product_id) VALUES (1, 2);`
	expected2 := []string{"SELECT * FROM users;", "INSERT INTO orders (user_id, product_id) VALUES (1, 2);"}
	result = SplitSQL(sql2)
	if !reflect.DeepEqual(result, expected2) {
		t.Errorf("SplitSQL() = %v, expected %v", result, expected2)
	}

	sql3 := `SELECT * FROM users; Insert 
	INTO orders (user_id, product_id) VALUES (1, 2);`
	expected3 := []string{"SELECT * FROM users;", `Insert 
	INTO orders (user_id, product_id) VALUES (1, 2);`}
	result = SplitSQL(sql3)
	if !reflect.DeepEqual(result, expected3) {
		t.Errorf("SplitSQL() = %v, expected %v", result, expected3)
	}

	sql4 := `SELECT * FROM users; Insert /* comment1
	comment2 */ INTO orders (user_id, product_id) VALUES (1, 2) ; 
	`
	expected4 := []string{"SELECT * FROM users;", `Insert /* comment1
	comment2 */ INTO orders (user_id, product_id) VALUES (1, 2) ;`}
	result = SplitSQL(sql4)
	if !reflect.DeepEqual(result, expected4) {
		t.Errorf("SplitSQL() = %v, expected %v", result, expected4)
	}

	sql5 := `SELECT * FROM users; Insert comment INTO orders (user_id, product_id) VALUES ("1;", 2) ;
	`
	expected5 := []string{"SELECT * FROM users;", `Insert comment INTO orders (user_id, product_id) VALUES ("1;", 2) ;`}
	result = SplitSQL(sql5)
	if !reflect.DeepEqual(result, expected5) {
		t.Errorf("SplitSQL() = %v, expected %v", result, expected5)
	}

}
