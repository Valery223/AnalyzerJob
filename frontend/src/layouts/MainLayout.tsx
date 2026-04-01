import { Link, Outlet } from 'react-router-dom';
import { useAuth } from '../context/AuthContext'; 


export default function MainLayout() {
   const { logout } = useAuth(); 
  return (
    <div className="flex h-screen bg-gray-100">
      {/* Боковая панель (Sidebar) */}
      <aside className="w-64 bg-white border-r shadow-sm flex flex-col">
        <div className="p-4 border-b">
          <h2 className="text-xl font-bold text-blue-600">AI Analyzer</h2>
        </div>
        <nav className="flex-1 p-4 flex flex-col gap-2">
          {/* Компонент Link из react-router-dom меняет URL без перезагрузки страницы */}
          <Link to="/" className="p-2 hover:bg-gray-50 rounded text-gray-700">Dashboard</Link>
          <Link to="/new" className="p-2 hover:bg-gray-50 rounded text-gray-700">New Analysis</Link>
        </nav>
      </aside>

      {/* Основная часть справа */}
      <div className="flex-1 flex flex-col">
        {/* Верхняя панель (Top Nav) */}
        <header className="h-16 bg-white border-b flex items-center justify-end px-6 shadow-sm">
          <div className="flex items-center gap-4">
            <span className="text-sm font-medium text-gray-700">user@example.com</span>
            <button onClick={logout} className="text-sm text-red-500 hover:underline">
              Выйти
            </button>
          </div>
        </header>

        {/* Контент страницы */}
        <main className="flex-1 overflow-auto">
          <Outlet />
        </main>
      </div>
    </div>
  );
}