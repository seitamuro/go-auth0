import logo from './logo.svg';
import './App.css';
import { useAuth0 } from '@auth0/auth0-react';
import { useEffect, useState } from 'react';

const API_URL = "http://localhost:8000";

const useAuth0Token = () => {
  const { isAuthenticated, user, getAccessTokenSilently } = useAuth0();
  const [accessToken, setAccessToken] = useState(null);

  useEffect(() => {
    const fetchToken = async () => {
      setAccessToken(await getAccessTokenSilently());
    };

    if(isAuthenticated) {
      fetchToken();
    }
  }, [isAuthenticated, user?.sub]);

  return accessToken;
}

function App() {
  const { loginWithRedirect, isAuthenticated } = useAuth0()
  const token = useAuth0Token();
  const [me, setMe] = useState(null);
  const [error, setError] = useState(null);

  const onClickLogin = () => {
    loginWithRedirect({
      authorizationParams: {
        redirect_uri: 'http://localhost:3000/'
      }
    })
  }

  const onClickCall = async () => {
    try {
      const res = await fetch(`${API_URL}/v1/users/me`, {
        method: "GET",
        mode: "cors",
        headers: {
          Authorization: `Bearer ${token}`,
          "Content-Type": "application/json",
        },
      });
      if(!res.ok) {
        throw new Error(res.statusText);
      }
      const me = await res.json();
      setError(null);
      setMe(me);
    } catch (error) {
      console.log("error", error);
      setError(error);
    }
  };

  return (
    <div className="App">
      <button onClick={onClickLogin} disabled={isAuthenticated}>
        {isAuthenticated ? "ログイン済み" : "ログイン"}
      </button>
      <div>
        <button onClick={onClickCall}>ユーザー情報を取得</button>
      </div>
      <div>
        <p>ユーザー: {JSON.stringify(me)}</p>
        <p>エラー: {error ? error.toString() : ""}</p>
      </div>
    </div>
  );
}

export default App;
