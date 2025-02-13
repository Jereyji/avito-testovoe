package queries

const (
	QueryFetchUserByID = `
		SELECT id, username, hashed_password, balance 
		FROM users 
		WHERE id = $1;
	`

	QueryFetchUserByUsername = `
		SELECT id, username, hashed_password, balance 
		FROM users 
		WHERE username = $1;
	`

	QueryCreateUser = `
		INSERT INTO users (id, username, hashed_password, balance)
		VALUES ($1, $2, $3, $4);
	`

	QueryUpdateUser = `
		UPDATE users
		SET username = $2, hashed_password = $3, balance = $4
		WHERE id = $1;
	`

	QueryDeleteUser = `
		DELETE FROM users
		WHERE id = $1;
	`
)

const (
	QueryFetchMerchByID = `
		SELECT id, name, price
		FROM merch
		WHERE id = $1;
	`

	QueryCreateMerch = `
		INSERT INTO merch (id, name, price)
		VALUES ($1, $2, $3);
	`

	QueryUpdateMerch = `
		UPDATE merch
		SET name = $2, price = $3
		WHERE id = $1;
	`

	QueryDeleteMerch = `
		DELETE FROM merch
		WHERE id = $1;
	`
)

const (
	QueryFetchTransactionsByUserID = `
		SELECT id, sender_id, recepient_id, amount, created_at 
		FROM transactions 
		WHERE sender_id = $1 OR recepient_id = $1;
	`

	QueryCreateTransaction = `
		INSERT INTO transactions (id, sender_id, recepient_id, amount, created_at)
		VALUES ($1, $2, $3, $4, $5);
	`

	QueryUpdateTransaction = `
		UPDATE transactions
		SET sender_id = $2, recepient_id = $3, amount = $4
		WHERE id = $1;
	`

	QueryDeleteTransaction = `
		DELETE FROM transactions
		WHERE id = $1;
	`
)

const (
	QueryFetchPurchasesByUserID = `
		SELECT id, user_id, merch_id, created_at
		FROM purchases
		WHERE user_id = $1;
	`

	QueryCreatePurchase = `
		INSERT INTO purchases (id, user_id, merch_id, created_at)
		VALUES ($1, $2, $3, NOW());
	`

	QueryUpdatePurchase = `
		UPDATE purchases
		SET user_id = $2, merch_id = $3
		WHERE id = $1;
	`

	QueryDeletePurchase = `
		DELETE FROM purchases
		WHERE id = $1;
	`
)
