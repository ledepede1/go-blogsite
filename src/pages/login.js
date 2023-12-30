import React from 'react';
import { useNavigate } from 'react-router-dom';

function LoginForm() {
const navigate = useNavigate();

async function handleLoginSubmit(event) { 
  event.preventDefault();

  const username = event.target.elements["username-login"].value;
  const password = event.target.elements["password-login"].value;

  // Checking if username and or password are empty
  if (username.trim().length === 0 || password.trim().length === 0) {
      alert("The form can't be blank");
      return;
  }

  // Send login data to the server
  const response = await fetch("http://localhost:8080/backend/login", {
      method: "POST",
      headers: {
          "Content-Type": "application/json",
      },
      body: JSON.stringify({ username, password }),
  });

  console.log(response.status)

  if (response.status === 200) {
      alert("Successfully logged in");
  } else if (response.status === 204) {
      alert("Either username or password wrong");
  } else {
      alert("Error, try again");
  }
}

function sendToSignUpPage() {
  navigate('/')
}
  

  return (
    <>
        <h1>Login</h1>
        <form onSubmit={handleLoginSubmit} className="login-form">
            <label>
            Username
            <input type="text" name="username-login" />
            </label>

            <label>
            Password
            <input type="password" name="password-login" />
            </label>
            
            <button type="submit">Login</button>
            <button
                type="button"
                onClick={sendToSignUpPage}
                className="login-button"
            >
                Dont have an account?
            </button>
        </form>
    </>
  );
}

export default LoginForm;