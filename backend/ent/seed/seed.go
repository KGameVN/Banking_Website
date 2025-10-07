package seed

import (
	"context"
	"log"
	"time"

	"comb.com/banking/ent"
)

func SeedData(ctx context.Context, client *ent.Client) {
	// Chạy 1 lần: nếu đã có user thì bỏ qua
	count, err := client.User.Query().Count(ctx)
	if err != nil {
		log.Fatalf("count users failed: %v", err)
	}
	if count > 0 {
		log.Println("Seed skipped: data already exists.")
		return
	}

	// --- Users ---
	u1, err := client.User.Create().
		SetUsername("User1").
		SetEmail("user@test.com").
		SetPassword("1234").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user1: %v", err)
	}

	u2, err := client.User.Create().
		SetUsername("User2").
		SetEmail("user2@test.com").
		SetPassword("1111").
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating user2: %v", err)
	}

	// --- Accounts ---
	acc1, err := client.UserAccount.Create().
		SetBalance(100).
		SetUser(u1).
		SetAccountNumber(1111). // thêm account number
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating account1: %v", err)
	}

	acc2, err := client.UserAccount.Create().
		SetBalance(1000).
		SetUser(u2).
		SetAccountNumber(2222). // thêm account number
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating account2: %v", err)
	}

	// --- Profile for User1 ---
	_, err = client.UserProfile.Create().
		SetFirstname("Nguyen").
		SetLastname("Van A").
		SetAddress("Hanoi").
		SetGender("male").
		SetBirthday("1990-01-01").
		SetUser(u1).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating profile: %v", err)
	}

	// --- Tokens ---
	_, err = client.Token.Create().
		SetToken("eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6IlVzZXIxIiwicGFzc3dvcmQiOiIxMjM0In0.PXp6WEHdDJXYHhpg_WwmlyRW_u-Y3zrajRn9729Zaxc").
		SetType("login").
		SetExpiredtime(time.Now().Add(24 * time.Hour)).
		SetUser(u1).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating token1: %v", err)
	}

	_, err = client.Token.Create().
		SetToken("token123").
		SetType("login").
		SetExpiredtime(time.Now().Add(24 * time.Hour)).
		SetUser(u2).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating token2: %v", err)
	}

	// --- Transfer (User1 -> User2, 50) ---
	_, err = client.Transfer.Create().
		SetAmount(50).
		SetFromAccount(acc1).
		SetToAccount(acc2).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating transfer: %v", err)
	}

	// --- Transactions (1-1 với account theo SQL) ---
	_, err = client.Transaction.Create().
		SetAmount(50).
		SetAccount(acc1).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating transaction1: %v", err)
	}

	_, err = client.Transaction.Create().
		SetAmount(200).
		SetAccount(acc2).
		Save(ctx)
	if err != nil {
		log.Fatalf("failed creating transaction2: %v", err)
	}

	log.Println("✅ Dummy data inserted successfully (first run only)!")
}
