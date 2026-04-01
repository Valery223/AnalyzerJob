import { useState } from 'react';
import { useNavigate } from 'react-router-dom';
import { api } from '../api/axios';

export default function NewAnalysisPage() {
  const navigate = useNavigate();

  // Состояния для полей формы
  const [title, setTitle] = useState('');
  const [company, setCompany] = useState('');
  const [description, setDescription] = useState('');
  
  // Состояния для процесса (чтобы блокировать кнопку)
  const [isLoading, setIsLoading] = useState(false);
  const [statusText, setStatusText] = useState(''); // Текст, чтобы юзер понимал, что происходит
  const [error, setError] = useState('');

  // Обработка отправки формы
  const handleSubmit = async (e: React.FormEvent) => {
    e.preventDefault();
    setIsLoading(true);
    setError('');

    try {
      // Шаг 1: Создаем саму вакансию в БД
      setStatusText('Сохраняем вакансию...');
      const createRes = await api.post('/vacancies/', {
        title,
        company,
        description
      });
      
      // Бэкенд возвращает созданный объект с id
      const newVacancyId = createRes.data.id;

      // Шаг 2: Запускем ИИ-генератор
      setStatusText('ИИ анализирует текст и придумывает вопросы (это может занять время)...');
      await api.post(`/vacancies/${newVacancyId}/generate`);

      // Шаг 3: Переходим на страницу этой вакансии
      navigate(`/vacancy/${newVacancyId}`);
      
    } catch (err: any) {
      console.error(err);
      setError(err.response?.data?.message || 'Произошла ошибка при обработке. Проверьте консоль.');
      setIsLoading(false);
    }
  };

  return (
    <div className="p-8 max-w-4xl mx-auto h-full flex flex-col">
      <h1 className="text-3xl font-bold text-gray-800 mb-2">Новый анализ вакансии</h1>
      <p className="text-gray-500 mb-8">Вставьте текст вакансии, и наш ИИ подготовит вас к собеседованию.</p>

      {error && (
        <div className="bg-red-100 text-red-700 p-4 rounded-lg mb-6 border border-red-200">
          {error}
        </div>
      )}

      <form onSubmit={handleSubmit} className="flex-1 flex flex-col gap-6">
        <div className="grid grid-cols-1 md:grid-cols-2 gap-6">
          {/* Поле: Должность (Target Role) */}
          <div>
            <label className="block text-sm font-semibold text-gray-700 mb-1">
              Желаемая должность (Target Role)
            </label>
            <input
              type="text"
              required
              placeholder="Например: Backend Dev (Golang)"
              value={title}
              onChange={(e) => setTitle(e.target.value)}
              disabled={isLoading}
              className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none disabled:bg-gray-100"
            />
          </div>

          {/* Поле: Компания */}
          <div>
            <label className="block text-sm font-semibold text-gray-700 mb-1">
              Компания
            </label>
            <input
              type="text"
              required
              placeholder="Например: Google, Yandex..."
              value={company}
              onChange={(e) => setCompany(e.target.value)}
              disabled={isLoading}
              className="w-full p-3 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none disabled:bg-gray-100"
            />
          </div>
        </div>

        {/* Большое текстовое поле для описания */}
        <div className="flex-1 flex flex-col">
          <label className="block text-sm font-semibold text-gray-700 mb-1">
            Текст вакансии (Job Description)
          </label>
          <textarea
            required
            placeholder="Вставьте сюда полное описание вакансии (требования, обязанности, условия)..."
            value={description}
            onChange={(e) => setDescription(e.target.value)}
            disabled={isLoading}
            className="w-full flex-1 p-4 border border-gray-300 rounded-lg focus:ring-2 focus:ring-blue-500 outline-none resize-none disabled:bg-gray-100 min-h-[300px]"
          />
        </div>

        {/*  КНОПКА */}
        <div className="mt-4 pb-8">
          <button
            type="submit"
            disabled={isLoading}
            className={`w-full py-5 text-xl font-bold rounded-xl transition-all shadow-lg 
              ${isLoading 
                ? 'bg-blue-400 text-white cursor-not-allowed animate-pulse' 
                : 'bg-blue-600 hover:bg-blue-700 text-white hover:shadow-xl hover:-translate-y-1'
              }`}
          >
            {isLoading ? statusText : 'СГЕНЕРИРОВАТЬ ВОПРОСЫ ДЛЯ ИНТЕРВЬЮ'}
          </button>
        </div>
      </form>
    </div>
  );
}