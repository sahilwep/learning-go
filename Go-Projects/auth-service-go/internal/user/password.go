/*

	- We will Never decrypts passwords, hashing is one-way process.
	- We only compare hashes
	- bcrypt is slower & it's prevent burteforce.

	// Bcrypt Depth:
		- let say we do:
			Password "hello" → hash X
			Password "hello" again → hash Y
			X ≠ Y
			Then how do we compare?

		- We do NOT compare hash == hash
		- We compare password -> hash -> match?

		// Deep Dive into how Bcrypt works:
			- let say for any random string "xyz":

				hash = $2a$10$N9qo8uLOickgx2ZMRZo5e.eFQz9f2e4UeFQz9f2e4UeFQz9f2e

			- That single hash string contains 4 things

				$2a$10$<salt><hashed-password>

				$2a$ → bcrypt algorithm version
				$10$ → cost factor (work factor)
				<salt> → RANDOM data (this is why hashes differ)
				<hash> → actual password hash
			- Salt embedded inside the hash itself.

			// Why same password gives different hashes:
				- Every time you call GenerateFromPassword:
					- bcrypt generates a new random salt
					- combines: slat + password
					- hashes them together
					- store salt + hash in one string.
					- so same string can have different salts
						- different salt => different design
							- This protects against Rainbow Table attack
							- Precomputed hash attack
							- Mass cred leaks.

			// How Bcrypt compares the passwords:
				- we will do:
					bcrypt.CompareHashAndPassword(storedHash, inputPassword)
				- Extracts:
					- salt from 'storedHash'
					- cost from 'storedHash'

				- Rehash with 'inputPassword' + 'extractedSalt'
				- compare it with {newlyComputeHash == storedHashHash}

				- Important: you never manually handel the salt, decrypt anything, compare hashes directly, bcrypt handel it safely

			// Why We never Decrypt password:
				- If password were decayable: insider threat = catastrophic, compliance failure (GDPR, SOC2, ISO).
				- Password must be One-way, Non-reversible, slow to compute, bcrypt satisfies all three.


			// Can We crack a password if DB is leaked & we got the storeHash:
				- Yes, in theory, a leaked bcrypt hash can be cracked.
				- In practice, bcrypt is designed to make it economically and time-wise infeasible.
				- let Say:
					- we got hash:
						- extract slat & cost
						- hash it with the list of password
						- Compare it with the stored hash & check if they were same?
							- If same ==> cracked..

				- In reality, There are multiple factors which restrict password to be cracked & time-wise infeasible.

				// ToolKit To protect password from cracked:
					- Slowness:
						- cost = 10 => 100–200 hashes per second per CPU core
						- if attacker with:
							- 1 GPU ≈ 10k-50k guesses/sec (optimistic), A large cluster ≈ millions/sec (very expensive)

					- Password Entropy:
						- Weak password => "hello"
						- Strong password = "r9!QmZ#eT2@Xk7"
						- Entropy ≈ 90+ bits, Brute forcing this would take thousands of years, even massive GPU frames..
						- This is why we have:
							- Minimum length
							- Complexity rules
							- Password Policy, & these things making it near to impossible.

					- Cost Factor:
						- Cost 10 => ok today		=>  100ms
						- Cost 12 => Stronger		=>  400ms
						- Cost 14 => Very strong	=>	1.6sec
						- For user login still OK, For attackers it's pain in A**
						- bcrypt let's you increase the cost in the future without changing passwords.

			// If bcrypt hashes can be brute forced offline, why is it secure ?
				- It's the tradeoff that we can calculate, like the time to crack if the hashes are leak & complexity of our hashes which is sorely depends on random salts & cost factors, so these things together makes the whole process more powerful.
				- This is just one layer of defense, it's just for damage control.

		- As of now bcrypt is one layer of defense, we will add more like MFA, Rate limiting, MFA, password length enforcement, breach detection, cred stuffing protection..



“If bcrypt hashes can be brute forced offline, why is it secure?”

*/

package user

import "golang.org/x/crypto/bcrypt"

/*
@HashPassword
  - hash the Incoming string into bcrypt
*/
func HashPassword(password string) (string, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return "", err
	}

	return string(hash), nil
}

/*
@CheckPassword
  - compare the given password & hash
*/
func CheckPassword(password, hash string) error {
	return bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
}
