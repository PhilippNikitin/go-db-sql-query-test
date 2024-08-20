package main

import (
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	_ "modernc.org/sqlite"
)

func Test_SelectClient_WhenOk(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NoError(t, err)
	}

	defer db.Close()

	clientID := 1

	cl, err := selectClient(db, clientID)

	require.NoError(t, err) // если ошибка не равна nil - завершаем тест

	assert.Equal(t, clientID, cl.ID) // проверяем, если поле ID полученного объекта Client совпадает с переменной clientID
	assert.NotEmpty(t, cl.FIO)       // проверяем, что поле FIO не пустое
	assert.NotEmpty(t, cl.Login)     // проверяем, что поле Login не пустое
	assert.NotEmpty(t, cl.Birthday)  // проверяем, что поле Birthday не пустое
	assert.NotEmpty(t, cl.Email)     // проверяем, что поле Email не пустое
}

func Test_SelectClient_WhenNoClient(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NoError(t, err)
	}

	defer db.Close()

	clientID := -1

	// напиши тест здесь
	cl, err := selectClient(db, clientID)

	require.Equal(t, sql.ErrNoRows, err) // проверяем, что есть ошибка, которую вернула функция selectClient, и она равна sql.ErrNoRows

	assert.Empty(t, cl.ID)       // проверяем, что поле ID у возвращенной структуры пустое
	assert.Empty(t, cl.FIO)      // проверяем, что поле FIO у возвращенной структуры пустое
	assert.Empty(t, cl.Login)    // проверяем, что поле Login у возвращенной структуры пустое
	assert.Empty(t, cl.Birthday) // проверяем, что поле Birthday у возвращенной структуры пустое
	assert.Empty(t, cl.Email)    // проверяем, что поле Email у возвращенной структуры пустое

}

func Test_InsertClient_ThenSelectAndCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NoError(t, err)
	}

	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	id, err := insertClient(db, cl)
	cl.ID = id
	require.NotEmpty(t, id) // проверяем, что вернулся не пустой идентификатор, иначе завершаем тест
	require.NoError(t, err) // проверяем, что вернулась ошибка, равная nil, иначе завершаем тест

	// получаем объект client по идентификатору id
	insertedCl, err := selectClient(db, id)

	require.NoError(t, err) // проверяем, что вернулась ошибка, равная nil, иначе завершаем тест

	assert.Equal(t, cl.ID, insertedCl.ID)             // проверяем, что идентификаторы клиента cl и полученного объекта insertedCl равны
	assert.Equal(t, cl.FIO, insertedCl.FIO)           // проверяем, что значения поля FIO клиента cl и полученного объекта insertedCl равны
	assert.Equal(t, cl.Login, insertedCl.Login)       // проверяем, что значения поля Login клиента cl и полученного объекта insertedCl равны
	assert.Equal(t, cl.Birthday, insertedCl.Birthday) // проверяем, что значения поля Birthday клиента cl и полученного объекта insertedCl равны
	assert.Equal(t, cl.Email, insertedCl.Email)       // проверяем, что значения поля Birthday клиента cl и полученного объекта insertedCl равны
}

func Test_InsertClient_DeleteClient_ThenCheck(t *testing.T) {
	// настройте подключение к БД
	db, err := sql.Open("sqlite", "demo.db")
	if err != nil {
		require.NoError(t, err)
	}

	defer db.Close()

	cl := Client{
		FIO:      "Test",
		Login:    "Test",
		Birthday: "19700101",
		Email:    "mail@mail.com",
	}

	// добавляем клиента в БД
	id, err := insertClient(db, cl)
	require.NotEmpty(t, id) // проверяем, что вернулся не пустой идентификатор, иначе завершаем тест
	require.NoError(t, err) // проверяем, что вернулась ошибка, равная nil, иначе завершаем тест

	// получаем клиента из БД при помощи функции selectClient
	_, err = selectClient(db, id)
	require.NoError(t, err) // если функция вернула ошибку, завершаем тест

	// удаляем клиента при помощи функции deleteClient
	err = deleteClient(db, id)
	require.NoError(t, err) // если функция вернула ошибку, завершаем тест

	_, err = selectClient(db, id)
	require.Equal(t, sql.ErrNoRows, err) // проверяем, что ошибка, которую вернула функция selectClient, равна sql.ErrNoRows

}
