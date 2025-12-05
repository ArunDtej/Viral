import { env } from '$env/dynamic/public';
import { browser } from '$app/environment';

export const API_BASE_URL = (browser && env.PUBLIC_API_URL) || "http://localhost:8080";
