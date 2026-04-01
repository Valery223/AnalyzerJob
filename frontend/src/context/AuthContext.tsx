import { createContext, useContext, useState, type ReactNode } from 'react';
import { useNavigate } from 'react-router-dom';

// контекст
interface AuthContextType {
  token: string | null;
  login: (newToken: string) => void;
  logout: () => void;
}

const AuthContext = createContext<AuthContextType | undefined>(undefined);

export function AuthProvider({ children }: { children: ReactNode }) {
  const navigate = useNavigate();
  // При загрузке приложения пытаемся достать токен из хранилища браузера
  const [token, setToken] = useState<string | null>(localStorage.getItem('token'));

  // Функция для успешного логина
  const login = (newToken: string) => {
    localStorage.setItem('token', newToken); // Сохраняем в браузер
    setToken(newToken);                      // Сохраняем в память React
    navigate('/');                           // Перекидываем на Дашборд
  };

  // Функция для выхода
  const logout = () => {
    localStorage.removeItem('token'); // Удаляем из браузера
    setToken(null);
    navigate('/login');               // Перекидываем на страницу входа
  };

  return (
    <AuthContext.Provider value={{ token, login, logout }}>
      {children}
    </AuthContext.Provider>
  );
}

// хук для использования контекста в компонентах
export const useAuth = () => {
  const context = useContext(AuthContext);
  if (!context) throw new Error("useAuth должен использоваться внутри AuthProvider");
  return context;
};