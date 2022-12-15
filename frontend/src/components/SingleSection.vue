<template>
  <base-card>
    <h3>{{ sectionTitle(section) }}</h3>
    <p>{{ section.title }}</p>
  </base-card>
</template>
<script>
import BaseCard from "@/components/UI/BaseCard.vue";

export default {
  components: {BaseCard},
  props: ["section"],
  methods: {
    sectionTitle(section) {
      return section.id + "@" + section.title;
    }
  }};
</script>