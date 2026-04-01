import { Navigate, Outlet } from 'react-router-dom';
import { useAuth } from '../context/AuthContext';

export default function ProtectedRoute() {
  const { token } = useAuth();

  // Если токена нет, принудительно перекидываем на логин 
  // Если есть, рисуем дочерние компоненты (Outlet)
  return token ? <Outlet /> : <Navigate to="/login" replace />;
}