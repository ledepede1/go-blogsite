import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Login from './pages/login';
import SignupForm from './pages/signup';
import './App.css';

function App() {
    return (
        <Router>
            <div className="App">
                <div className="login-container">
                    <Routes>
                        <Route path="/login" element={<Login />} />
                        <Route path="/" element={<SignupForm />} />
                    </Routes>
                </div>
            </div>
        </Router>
    );
}

export default App;