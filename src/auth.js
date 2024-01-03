import Cookies from 'js-cookie';
import { useHistory } from 'react-router-dom';

export const getAccessToken = () => Cookies.get('access_token');
export const getRefreshToken = () => Cookies.get('refresh_token');
export const isAuthenticated = () => !!getAccessToken();

export const authenticate = async () => {
  if (getRefreshToken()) {
    try {
      const tokens = await refreshToken(); // Assuming refreshToken is an async function that fetches tokens

      const expires = (tokens.expires_in || 60 * 60) * 1000;
      const inOneHour = new Date(new Date().getTime() + expires);

      Cookies.set('access_token', tokens.access_token, { expires: inOneHour });
      Cookies.set('refresh_token', tokens.refresh_token);

      return true;
    } catch (error) {
      redirectToLogin();
      return false;
    }
  }

  redirectToLogin();
  return false;
};

const refreshToken = async () => {
  // Implement the logic to refresh the token, e.g., make an API call
  // Return the new tokens
};

const redirectToLogin = () => {
  //const history = useHistory();
  //history.push('/login');
};

export default { authenticate, getAccessToken, getRefreshToken, isAuthenticated };
