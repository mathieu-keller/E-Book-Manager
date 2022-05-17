const API_PREFIX = '/api';
export const DOWNLOAD_API = (id: number) => `${API_PREFIX}/book/download/${id}`;
export const DOWNLOAD_ORIGINAL_API = (id: number) => `${API_PREFIX}/book/original/download/${id}`;
export const SEARCH_API = (search: string, page: number) => `${API_PREFIX}/book?q=${encodeURIComponent(search)}&page=${page}`;
export const BOOK_API = (title: string) => `${API_PREFIX}/book/${title}`;
export const LIBRARY_API = (page: number) => `${API_PREFIX}/library/all?page=${page}`;
export const COLLECTION_API = (title: string) => `${API_PREFIX}/collection?title=${title}`;
export const UPLOAD_API = `${API_PREFIX}/upload/multi`;
