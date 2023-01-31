import logo from './logo.svg';
import './App.css';
import { useAuth0 } from '@auth0/auth0-react';

function App() {
  const { loginWithRedirect, isAuthenticated } = useAuth0

  const onClickLogin = () => {
    loginWithRedirect()
  }

  return (
    <div className="App">
      <button onClick={onClickLogin} disabled={isAuthenticated}>
        {isAuthenticated ? "ログイン済み" : "ログイン"}
      </button>
    </div>
  );
}

export default App;
