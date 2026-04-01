import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { AuthProvider } from './context/AuthContext'; // Импортируем провайдер
import ProtectedRoute from './components/ProtectedRoute'; // Импортируем защиту

import MainLayout from './layouts/MainLayout';
import DashboardPage from './pages/DashboardPage';
import LoginPage from './pages/LoginPage';
import NewAnalysisPage from './pages/NewAnalysisPage';
import VacancyDetailPage from './pages/VacancyDetailPage';

function App() {
  return (
    <BrowserRouter>
      {/* AuthProvider дает доступ к токену компонентам внутри него */}
      <AuthProvider>
        <Routes>
          {/* Публичный маршрут */}
          <Route path="/login" element={<LoginPage />} />

          {/* Защищенные маршруты */}
          {/* Сначала проверяем токен (ProtectedRoute), потом рисуем Layout, потом саму страницу */}
          <Route element={<ProtectedRoute />}>
            <Route element={<MainLayout />}>
              <Route path="/" element={<DashboardPage />} />
              <Route path="/new" element={<NewAnalysisPage />} />
              <Route path="/vacancy/:id" element={<VacancyDetailPage />} />
            </Route>
          </Route>
        </Routes>
      </AuthProvider>
    </BrowserRouter>
  );
}

export default App;