import React from 'react';
import { BrowserRouter as Router, Route, Routes } from 'react-router-dom';
import Login from './pages/login';
import SignupForm from './pages/signup';
import MainPage from './pages/mainpage';
import { Navigate } from 'react-router-dom';
import './App.css';


export const PrivateRoute = ({ children}) => {
    const isAuthenticated = true;
        
    if (isAuthenticated ) {
      return children
    }
      
    return <Navigate to="/login" />
}

function App() {

    return (
        <Router>
            <div className="App">
                <div className="login-container">
                    <Routes>
                        <Route path="/login" element={<Login />} />
                        <Route path="/" element={<SignupForm />} />
                        <Route 
                        path="/main" 
                        element={
                        <PrivateRoute>
                        <MainPage />
                        </PrivateRoute>
                        } 
                        />
                    </Routes>
                </div>
            </div>
        </Router>
    );
}

export default App;