<template>
  <div>
    <h1>Online Auctions</h1>
    <div>
      <input v-model="title" placeholder="Auction Title" />
      <input v-model="description" placeholder="Description" />
      <input v-model.number="startingBid" type="number" placeholder="Starting Bid" />
      <input v-model="endTime" type="datetime-local" placeholder="End Time" />
      <button @click="createAuction">Create Auction</button>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      title: '',
      description: '',
      startingBid: 0,
      endTime: ''
    };
  },
  methods: {
    async createAuction() {
      try {
        const response = await axios.post('http://localhost:8080/api/auctions', {
          title: this.title,
          description: this.description,
          startingBid: this.startingBid,
          endTime: this.endTime
        });
        console.log('Auction created:', response.data);
      } catch (error) {
        console.error('Error creating auction:', error);
      }
    }
  }
};
</script>
