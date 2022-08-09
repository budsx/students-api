# Students API

---

**API for students**

- `GET` /students
- `GET` /students/{id}
- `POST` /students
- `PUT` /students/{id}
- `DELETE` /students/{id}

**Register & Login**

- `POST` /auth/register

```
{
    "name": "Budi",
    "username": "budisetiawan",
    "password": "123456789"
}

```

**Register Success**

```
{
    "message": "User registered successfully"
}

```

**Register Failed**

```
{
    "message": "User already exists"
}

```

---

- `POST` /auth/login

```
{
    "username": "budisetiawan",
    "password": "123456789"
}

```

**Login Success**

```
{
    "access_token": {
        "access_token": "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6ImJ1ZGl3In0.XNdpE8HmQxKlRG2GBF2IZ0h8UbTtiH91siS5OaFHyEM"
    }
}

```

**Login Failed**

```
{
    "message": "Username and password do not match"
}

```
