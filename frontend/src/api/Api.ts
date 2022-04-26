const API_PREFIX = '/api';
export const DOWNLOAD_API = (id: number) => `${API_PREFIX}/download/${id}`;
export const UPLOAD_API = `${API_PREFIX}/upload/multi`;
export const SEARCH_API = (search: string, page: number) => `${API_PREFIX}/book?q=${encodeURIComponent(search)}&page=${page}`;
export const BOOK_API = (title: string) => `${API_PREFIX}/book/${title}`;
export const LIBRARY_API = (page: number) => `${API_PREFIX}/all?page=${page}`;
export const COLLECTION_API = (title: string) => `${API_PREFIX}/collection?title=${title}`;


