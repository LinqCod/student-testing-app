<template>
  <div class="section-list">
    <single-section v-for="section in sections" :key="section.id" :section="section">
    </single-section>
  </div>
</template>
<script>
import SingleSection from "./SingleSection.vue";
export default {
  components: { SingleSection },
  props: ["sections"]
};
</script>
<style></style>