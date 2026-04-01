import { useState, useEffect } from 'react';
import { Link } from 'react-router-dom';
import { api } from '../api/axios';
import type { Vacancy } from '../types';

export default function DashboardPage() {
  // Состояния (State)
  const [vacancies, setVacancies] = useState<Vacancy[]>([]);
  const [search, setSearch] = useState('');
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState('');

  // Функция загрузки данных с бэкенда
  const fetchVacancies = async (searchQuery: string = '') => {
    setIsLoading(true);
    setError('');
    try {
      // Делаем GET запрос. axios сам подставит токен из перехватчика!
      const response = await api.get('/vacancies/', {
        params: { search: searchQuery } // axios сам соберет ?search=...
      });
      // Если бэкенд возвращает null при пустом списке, защитимся от этого:
      setVacancies(response.data || []); 
    } catch (err: any) {
      setError('Не удалось загрузить вакансии. Сервер недоступен.');
      console.error(err);
    } finally {
      setIsLoading(false); // Выключаем загрузку в любом случае
    }
  };

  // useEffect (с пустым массивом []) срабатывает ровно ОДИН раз при открытии страницы.
  // Это аналог функции init() в Go.
  useEffect(() => {
    fetchVacancies();
  }, []);

  // Обработчик формы поиска
  const handleSearch = (e: React.FormEvent) => {
    e.preventDefault();
    fetchVacancies(search); // Ищем по введенному тексту
  };

  return (
    <div className="p-8 max-w-7xl mx-auto flex flex-col h-full">
      
      {/* Верхний блок: Заголовок и Кнопка создания */}
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold text-gray-800">Мои анализы</h1>
        <Link 
          to="/new" 
          className="bg-blue-600 hover:bg-blue-700 text-white px-6 py-2 rounded-lg font-medium transition-colors shadow-sm"
        >
          + Новый анализ ИИ
        </Link>
      </div>

      {/* Блок поиска */}
      <form onSubmit={handleSearch} className="mb-8 flex gap-2">
        <input
          type="text"
          placeholder="Поиск по названию или компании..."
          value={search}
          onChange={(e) => setSearch(e.target.value)}
          className="flex-1 max-w-md p-2 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none"
        />
        <button type="submit" className="bg-gray-800 text-white px-4 py-2 rounded-lg hover:bg-gray-700">
          Найти
        </button>
        {search && (
          <button 
            type="button" 
            onClick={() => { setSearch(''); fetchVacancies(''); }} 
            className="text-gray-500 hover:underline px-2"
          >
            Сбросить
          </button>
        )}
      </form>

      {/* Обработка состояний: Загрузка, Ошибка, Пустой список */}
      {isLoading ? (
        <div className="flex-1 flex justify-center items-center text-gray-500">
          Загрузка вакансий...
        </div>
      ) : error ? (
        <div className="bg-red-100 text-red-700 p-4 rounded-lg">{error}</div>
      ) : vacancies.length === 0 ? (
        <div className="flex-1 flex flex-col justify-center items-center text-gray-500 border-2 border-dashed border-gray-300 rounded-xl p-8">
          <p className="text-lg">У вас пока нет сохраненных вакансий.</p>
          <p className="text-sm">Нажмите "Новый анализ ИИ", чтобы начать.</p>
        </div>
      ) : (
        /* Сетка карточек вакансий */
        <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
          {vacancies.map((vacancy) => (
            <div key={vacancy.id} className="bg-white border rounded-xl p-6 shadow-sm hover:shadow-md transition-shadow flex flex-col">
              <div className="flex justify-between items-start mb-4">
                <div>
                  <h3 className="text-xl font-semibold text-gray-800">{vacancy.title}</h3>
                  <p className="text-blue-600 font-medium">{vacancy.company}</p>
                </div>
              </div>
              
              <p className="text-gray-600 text-sm line-clamp-3 mb-4 flex-1">
                {vacancy.description}
              </p>
              
              <div className="pt-4 border-t mt-auto flex justify-between items-center">
                <span className="text-xs text-gray-400">
                  {new Date(vacancy.created_at).toLocaleDateString('ru-RU')}
                </span>
                <Link 
                  to={`/vacancy/${vacancy.id}`} 
                  className="text-sm font-medium text-blue-600 hover:underline"
                >
                  Смотреть вопросы →
                </Link>
              </div>
            </div>
          ))}
        </div>
      )}

    </div>
  );
}