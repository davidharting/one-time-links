# One-time links

Securely share one-time links.

## Implementation

This is the tech I was excited to try out for this project!

- Uses go:embed to make templates available at runtime in the binary
- Encrypts / decrypts messages using https://github.com/ProtonMail/gopenpgp
- Stores messages in DynamoDB

## Work left

- Generate random password when a message is created
- Delete message after rendering
- Host this somewhere!
- Add favicon
