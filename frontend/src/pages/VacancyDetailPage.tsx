import { useState, useEffect } from 'react';
import { useParams, Link, useNavigate } from 'react-router-dom';
import { api } from '../api/axios';
import { type Vacancy } from '../types';

export default function VacancyDetailPage() {
  const { id } = useParams<{ id: string }>(); 
  const navigate = useNavigate();

  // Состояния
  const [vacancy, setVacancy] = useState<Vacancy | null>(null);
  const [isLoading, setIsLoading] = useState(true);
  const [error, setError] = useState('');
  
  // Состояние для чекбоксов
  const [learnedQuestions, setLearnedQuestions] = useState<Set<number>>(new Set());

  // Загрузка данных вакансии
  useEffect(() => {
    const fetchVacancy = async () => {
      try {
        const response = await api.get(`/vacancies/${id}`);
        setVacancy(response.data);
      } catch (err: any) {
        setError('Не удалось загрузить данные вакансии. Возможно, она была удалена.');
        console.error(err);
      } finally {
        setIsLoading(false);
      }
    };

    if (id) {
      fetchVacancy();
    }
  }, [id]);

  // Обработчик удаления вакансии
  const handleDelete = async () => {
    if (!window.confirm('Вы уверены, что хотите удалить этот анализ?')) return;
    
    try {
      await api.delete(`/vacancies/${id}`);
      navigate('/');
    } catch (err) {
      alert('Ошибка при удалении');
    }
  };

  // Обработчик клика по чекбоксу
  const toggleLearned = (index: number) => {
    setLearnedQuestions(prev => {
      const newSet = new Set(prev);
      if (newSet.has(index)) {
        newSet.delete(index); // Снимаем галку
      } else {
        newSet.add(index);    // Ставим галку
      }
      return newSet;
    });
  };

  // 1. Состояние: Идет загрузка
  if (isLoading) {
    return (
      <div className="flex h-full items-center justify-center">
        <div className="text-xl text-gray-500 animate-pulse">Загрузка результатов ИИ...</div>
      </div>
    );
  }

  // 2. Состояние: Ошибка
  if (error || !vacancy) {
    return (
      <div className="p-8 max-w-4xl mx-auto text-center">
        <div className="bg-red-100 text-red-700 p-6 rounded-xl mb-4 text-lg">{error}</div>
        <Link to="/" className="text-blue-600 hover:underline">← Вернуться на Дашборд</Link>
      </div>
    );
  }

  // 3. Состояние: Успех
  return (
    <div className="p-8 max-w-4xl mx-auto flex flex-col min-h-full">
      
      {/* Шапка: Кнопка назад и Заголовок */}
      <div className="mb-8">
        <Link to="/" className="text-gray-500 hover:text-gray-800 flex items-center gap-2 mb-6 w-fit transition-colors">
          <span className="text-xl">←</span> Вернуться к списку (Back to Dashboard)
        </Link>
        
        <div className="flex justify-between items-start gap-4">
          <div>
            <h1 className="text-4xl font-bold text-gray-900 mb-2">{vacancy.title}</h1>
            <p className="text-xl text-blue-600 font-medium">{vacancy.company}</p>
          </div>
          
          <button 
            onClick={handleDelete}
            className="px-4 py-2 text-red-600 border border-red-200 rounded-lg hover:bg-red-50 transition-colors"
          >
            Удалить анализ
          </button>
        </div>
      </div>

      {/* Описание вакансии (Свернутое) */}
      <div className="bg-white rounded-xl shadow-sm border p-6 mb-8">
        <h2 className="text-sm font-bold text-gray-400 uppercase tracking-wider mb-3">
          Исходный текст вакансии
        </h2>
        <div className="text-gray-600 text-sm whitespace-pre-wrap max-h-32 overflow-y-auto pr-2 custom-scrollbar">
          {vacancy.description}
        </div>
      </div>

      {/* Главный блок: Вопросы для собеседования */}
      <div className="flex-1">
        <h2 className="text-2xl font-bold text-gray-800 mb-6 flex items-center gap-3">
          Вопросы для подготовки
          <span className="bg-blue-100 text-blue-800 text-sm py-1 px-3 rounded-full">
            {vacancy.ai_questions?.length || 0} шт.
          </span>
        </h2>

        {!vacancy.ai_questions || vacancy.ai_questions.length === 0 ? (
          <div className="text-gray-500 italic p-6 bg-gray-50 rounded-xl border border-dashed text-center">
            ИИ не смог сгенерировать вопросы для этой вакансии.
          </div>
        ) : (
          <div className="space-y-4 pb-8">
            {vacancy.ai_questions.map((question, index) => {
              const isLearned = learnedQuestions.has(index);
              
              return (
                <div 
                  key={index}
                  className={`flex items-start gap-4 p-5 rounded-xl border transition-all duration-200
                    ${isLearned 
                      ? 'bg-green-50 border-green-200 opacity-75' 
                      : 'bg-white shadow-sm hover:shadow-md border-gray-200'
                    }`}
                >
                  {/*  чекбокс */}
                  <div className="pt-1">
                    <input 
                      type="checkbox"
                      checked={isLearned}
                      onChange={() => toggleLearned(index)}
                      className="w-6 h-6 rounded border-gray-300 text-green-600 focus:ring-green-500 cursor-pointer"
                    />
                  </div>
                  
                  {/* Текст вопроса */}
                  <div className="flex-1">
                    <p className={`text-lg ${isLearned ? 'text-green-800 line-through' : 'text-gray-800 font-medium'}`}>
                      {question}
                    </p>
                    {isLearned && (
                      <span className="text-xs font-bold text-green-600 uppercase tracking-wider mt-2 block">
                        Изучено (Learned) ✓
                      </span>
                    )}
                  </div>
                </div>
              );
            })}
          </div>
        )}
      </div>
      
    </div>
  );
}