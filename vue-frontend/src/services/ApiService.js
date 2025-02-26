import axios from 'axios';

const apiClient = axios.create({
    baseURL: process.env.VUE_APP_API_BASE_URL || 'http://localhost:8080/api', // Replace with your Go backend URL
    withCredentials: true // Important for cookies/sessions if used
});

// Add interceptor for JWT if you use it
apiClient.interceptors.request.use(
    (config) => {
        const token = localStorage.getItem('token'); // Or however you store your token
        if (token) {
            config.headers.Authorization = `Bearer ${token}`;
        }
        return config;
    },
    (error) => {
        return Promise.reject(error);
    }
);


const ApiService = {
    getProducts() {
        return apiClient.get('/products');
    },
    getProduct(id) {
        return apiClient.get(`/products/${id}`);
    },
    getAuctions() {
        return apiClient.get('/auctions');
    },
    getAuction(id) {
        return apiClient.get(`/auctions/${id}`);
    },
    getBids(auctionId) {
        return apiClient.get(`/bids?auction_id=${auctionId}`);
    },
    placeBid(auctionId, amount) {
        return apiClient.post('/bids', { auction_id: auctionId, amount });
    },
    register(user) {
        return apiClient.post('/register', user);
    },
    login(user) {
        return apiClient.post('/login', user);
    },
    // ... other API calls as needed (payments, admin, etc.)
};

export default ApiService;
