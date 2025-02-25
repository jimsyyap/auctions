<template>
  <div>
    <h1>Register</h1>
    <form @submit.prevent="register">
      <input v-model="username" type="text" placeholder="Username" required />
      <input v-model="email" type="email" placeholder="Email" required />
      <input v-model="password" type="password" placeholder="Password" required />
      <button type="submit">Register</button>
    </form>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      username: '',
      email: '',
      password: '',
    };
  },
  methods: {
    async register() {
      try {
        const response = await axios.post('http://localhost:8080/api/register', {
          username: this.username,
          email: this.email,
          password: this.password,
          role: 'Buyer',
        });
        alert(response.data.message);
      } catch (error) {
        alert('Registration failed');
        console.error(error);
      }
    },
  },
};
</script>
