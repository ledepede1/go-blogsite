import React from 'react';
import { useNavigate } from 'react-router-dom';

function SignupForm() {
    const navigate = useNavigate();

    function sendToLoginPage() {
        // Navigate to the "/login" route
        navigate('/login');
    }

    async function handleSignupSubmit(event) {
        event.preventDefault();

        const username = event.target.elements["username-signup"].value;
        const password = event.target.elements["password-signup"].value;

        // Checking if username and or password are empty
        if (username.trim().length === 0 || password.trim().length === 0) {
            alert("The form can't be blank");
            return;
        }

        // Send signup data to the server
        const response = await fetch("http://localhost:8080/backend/signup", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({ username, password }),
        });

        if (response.ok) {
            alert("Successfully created a new account");
        } else if (response.status === 423) {
            alert("Username already exists");
        } else {
            alert("Error, try again");
        }
    }

    return (
        <>
            <h1>Signup</h1>
            <form onSubmit={handleSignupSubmit} className="login-form">
                <label>
                Username
                <input type="text" name="username-signup" />
                </label>

                <label>
                Password
                <input type="password" name="password-signup" />
                </label>
                
                <button type="submit">Signup</button>
                <button
                    type="button"
                    onClick={sendToLoginPage}
                    className="login-button"
                >
                    Already have an account?
                </button>
            </form>
        </>
    );
}

export default SignupForm;