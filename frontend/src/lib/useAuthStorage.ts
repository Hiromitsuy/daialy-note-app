import { useCallback, useEffect, useState } from 'react';

const STORAGE_EVENT_KEY = 'auth-token-changed';

export default function useAuthStorage() {
  const getStorage = () => {
    if (typeof window === 'undefined') return null;
    return sessionStorage.getItem('token');
  };

  const [token, setTokenState] = useState(getStorage);

  useEffect(() => {
    const handleStorageChange = () => {
      setTokenState(getStorage());
    };

    window.addEventListener(STORAGE_EVENT_KEY, handleStorageChange);

    return () => {
      window.removeEventListener(STORAGE_EVENT_KEY, handleStorageChange);
    };
  }, []);

  const setToken = useCallback((newToken: string | null) => {
    if (typeof window === 'undefined') return;

    if (newToken) {
      sessionStorage.setItem('token', newToken);
    } else {
      sessionStorage.removeItem('token');
    }

    setTokenState(newToken);
    window.dispatchEvent(new Event(STORAGE_EVENT_KEY));
  }, []);

  return { token, setToken };
}
