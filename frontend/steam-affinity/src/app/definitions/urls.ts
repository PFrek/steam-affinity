export const baseURL = process.env.NODE_ENV === "production"
	? process.env.NEXT_PUBLIC_API_URL
	: "http://localhost:8080/api/v1/";
