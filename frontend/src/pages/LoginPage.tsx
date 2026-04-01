import { useState } from 'react';
import { api } from '../api/axios';
import { useAuth } from '../context/AuthContext';

export default function LoginPage() {
  const { login } = useAuth();
  
  // Состояния для полей ввода
  const [email, setEmail] = useState('');
  const [password, setPassword] = useState('');
  
  // Состояния для UI
  const [isRegisterMode, setIsRegisterMode] = useState(false); 
  const [error, setError] = useState('');
  const [isLoading, setIsLoading] = useState(false);

  // Обработчик нажатия на кнопку "Войти/Зарегистрироваться"
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault(); 
    setError('');
    setIsLoading(true);

    try {
      if (isRegisterMode) {
        // Запрос на регистрацию
        await api.post('/auth/register', { email, password });
        // Сразу после регистрации делаем логин, чтобы получить токен
        const res = await api.post('/auth/login', { email, password });
        login(res.data.token);
      } else {
        // Запрос на логин 
        const res = await api.post('/auth/login', { email, password });
        login(res.data.token); // Вызываем функцию из AuthContext
      }
    } catch (err: any) {
      setError(err.response?.data?.message || 'Произошла ошибка. Проверьте данные.');
    } finally {
      setIsLoading(false);
    }
  };

  return (
    <div className="flex h-screen items-center justify-center bg-gray-50">
      <div className="w-full max-w-md p-8 space-y-6 bg-white rounded-lg shadow-md">
        <h1 className="text-2xl font-bold text-center">
          {isRegisterMode ? 'Создать аккаунт' : 'Войти в систему'}
        </h1>

        {error && <div className="p-3 text-sm text-red-700 bg-red-100 rounded">{error}</div>}

        <form onSubmit={handleSubmit} className="space-y-4">
          <div>
            <label className="block text-sm font-medium text-gray-700">Email</label>
            <input
              type="email"
              required
              value={email}
              onChange={(e) => setEmail(e.target.value)}
              className="mt-1 w-full p-2 border border-gray-300 rounded focus:ring-blue-500 focus:border-blue-500"
            />
          </div>
          <div>
            <label className="block text-sm font-medium text-gray-700">Пароль</label>
            <input
              type="password"
              required
              value={password}
              onChange={(e) => setPassword(e.target.value)}
              className="mt-1 w-full p-2 border border-gray-300 rounded focus:ring-blue-500 focus:border-blue-500"
            />
          </div>

          <button
            type="submit"
            disabled={isLoading}
            className="w-full py-2 px-4 bg-blue-600 text-white rounded hover:bg-blue-700 disabled:bg-blue-300 transition-colors"
          >
            {isLoading ? 'Загрузка...' : (isRegisterMode ? 'Зарегистрироваться' : 'Войти')}
          </button>
        </form>

        <div className="text-center text-sm">
          <span className="text-gray-600">
            {isRegisterMode ? 'Уже есть аккаунт?' : 'Нет аккаунта?'}
          </span>{' '}
          <button
            type="button"
            onClick={() => setIsRegisterMode(!isRegisterMode)}
            className="text-blue-600 hover:underline"
          >
            {isRegisterMode ? 'Войти' : 'Зарегистрироваться'}
          </button>
        </div>
      </div>
    </div>
  );
}