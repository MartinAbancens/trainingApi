package handler

import (
	"testing"
)

func testGetCurrencyByIDgo(t *testing.T) {
	// db, mock, err := sqlmock.New()
	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	// }
	// defer db.Close()

	// mock.ExpectBegin()
	// mock.ExpectExec("UPDATE products").WillReturnResult(sqlmock.NewResult(1, 1))
	// mock.ExpectExec("INSERT INTO product_viewers").WithArgs(2, 3).WillReturnResult(sqlmock.NewResult(1, 1))
	// mock.ExpectCommit()

	// // now we execute our method
	// if err = recordStats(db, 2, 3); err != nil {
	// 	t.Errorf("error was not expected while updating stats: %s", err)
	// }

	// // we make sure that all expectations were met
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// }

	// -----------------------------------------

	// db, mock, err := sqlmock.New()
	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	// }
	// defer db.Close()

	// // create app with mocked db, request and response to test
	// app := &api{db}
	// req, err := http.NewRequest("GET", "http://localhost/posts", nil)
	// if err != nil {
	// 	t.Fatalf("an error '%s' was not expected while creating request", err)
	// }
	// w := httptest.NewRecorder()

	// // before we actually execute our api function, we need to expect required DB actions
	// rows := sqlmock.NewRows([]string{"id", "title", "body"}).
	// 	AddRow(1, "post 1", "hello").
	// 	AddRow(2, "post 2", "world")

	// mock.ExpectQuery("^SELECT (.+) FROM posts$").WillReturnRows(rows)

	// // now we execute our request
	// app.posts(w, req)

	// if w.Code != 200 {
	// 	t.Fatalf("expected status code to be 200, but got: %d", w.Code)
	// }

	// data := struct {
	// 	Posts []*post
	// }{Posts: []*post{
	// 	{ID: 1, Title: "post 1", Body: "hello"},
	// 	{ID: 2, Title: "post 2", Body: "world"},
	// }}
	// app.assertJSON(w.Body.Bytes(), data, t)

	// // we make sure that all expectations were met
	// if err := mock.ExpectationsWereMet(); err != nil {
	// 	t.Errorf("there were unfulfilled expectations: %s", err)
	// }
}
