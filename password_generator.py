import random
import string

def generate_password(length=12):
    """Generate a random password with letters, digits, and special characters."""
    characters = string.ascii_letters + string.digits + string.punctuation
    password = ''.join(random.choice(characters) for _ in range(length))
    return password

if __name__ == "__main__":
    try:
        length = int(input("Enter password length: "))
        if length < 6:
            print("Password should be at least 6 characters long.")
        else:
            print("Generated Password:", generate_password(length))
    except ValueError:
        print("Please enter a valid number.")
