export interface Vacancy {
  id: string;
  title: string;
  company: string;
  description: string;
  ai_questions: string[] | null;
  created_at: string;
}