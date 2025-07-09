const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080/api/v1';

export interface Student {
  id: number;
  student_code: string;
  first_name: string;
  last_name: string;
  email: string;
  phone?: string;
  date_of_birth?: string;
  address?: string;
  major?: string;
  year?: number;
  gpa?: number;
  status: 'active' | 'inactive' | 'graduated' | 'suspended';
  created_at: string;
  updated_at: string;
}

export interface ApiResponse<T> {
  data: T;
  count?: number;
  status?: string;
  message?: string;
  major?: string;
}

class ApiClient {
  private baseURL: string;

  constructor(baseURL: string) {
    this.baseURL = baseURL;
  }

  private async request<T>(endpoint: string, options?: RequestInit): Promise<T> {
    const url = `${this.baseURL}${endpoint}`;

    console.log(`Making ${options?.method || 'GET'} request to:`, url);

    const response = await fetch(url, {
      headers: {
        'Content-Type': 'application/json',
        ...options?.headers,
      },
      ...options,
    });

    console.log(`Response status: ${response.status}`);

    if (!response.ok) {
      const errorText = await response.text();
      console.error(`API Error: ${response.status} - ${errorText}`);
      throw new Error(`HTTP error! status: ${response.status} - ${errorText}`);
    }

    const data = await response.json();
    console.log('Response data:', data);
    return data;
  }

  // Student endpoints
  async getStudents(): Promise<ApiResponse<Student[]>> {
    return this.request<ApiResponse<Student[]>>('/students');
  }

  async getStudent(id: number): Promise<ApiResponse<Student>> {
    return this.request<ApiResponse<Student>>(`/students/${id}`);
  }

  async createStudent(studentData: Omit<Student, 'id' | 'created_at' | 'updated_at'>): Promise<ApiResponse<Student>> {
    return this.request<ApiResponse<Student>>('/students', {
      method: 'POST',
      body: JSON.stringify(studentData),
    });
  }

  async updateStudent(id: number, studentData: Partial<Student>): Promise<ApiResponse<Student>> {
    return this.request<ApiResponse<Student>>(`/students/${id}`, {
      method: 'PUT',
      body: JSON.stringify(studentData),
    });
  }

  async deleteStudent(id: number): Promise<{ message: string }> {
    return this.request<{ message: string }>(`/students/${id}`, {
      method: 'DELETE',
    });
  }

  async getStudentsByMajor(major: string): Promise<ApiResponse<Student[]>> {
    return this.request<ApiResponse<Student[]>>(`/students/major?major=${major}`);
  }

  async getStudentsByStatus(status: string = 'active'): Promise<ApiResponse<Student[]>> {
    return this.request<ApiResponse<Student[]>>(`/students/status?status=${status}`);
  }

  // Health check
  async healthCheck(): Promise<{ status: string }> {
    const url = `${this.baseURL.replace('/api/v1', '')}/health`;
    const response = await fetch(url);
    return response.json();
  }
}

export const apiClient = new ApiClient(API_BASE_URL);
