<template>
  <div>
    <h1>Auction Details</h1>
    <div v-if="auction">
      <h2>{{ auction.product.name }}</h2>
      <p>{{ auction.product.description }}</p>
      <p>Current Bid: {{ auction.current_bid }}</p>
      </div>
    <div v-else>
      Loading...
    </div>
    <div v-if="error">
      {{ error }}
    </div>
  </div>
</template>

<script>
import ApiService from '@/services/ApiService'; // Adjust path if needed

export default {
  data() {
    return {
      auction: null,
      error: null
    };
  },
  mounted() {
    this.fetchAuctionDetails();
  },
  methods: {
    async fetchAuctionDetails() {
      try {
        const auctionId = this.$route.params.id; // Get auction ID from route
        const response = await ApiService.getAuction(auctionId);
        this.auction = response.data;
        // Fetch product details separately if needed
        const productResponse = await ApiService.getProduct(this.auction.product_id); // Assuming product_id is available
        this.auction.product = productResponse.data;
      } catch (err) {
        console.error(err);
        this.error = err.message || "Error fetching auction details.";
      }
    }
  }
};
</script>
