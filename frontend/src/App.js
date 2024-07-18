import React, { useState } from 'react';
import Login from './components/Login';
import Dashboard from './components/Dashboard';

const App = () => {
    const [token, setToken] = useState(null);

    return (
        <div className="app-container">
            {token ? <Dashboard token={token} /> : <Login setToken={setToken} />}
        </div>
    );
};

export default App;
