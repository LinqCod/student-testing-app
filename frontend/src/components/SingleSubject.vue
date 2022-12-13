<template>
  <base-card>
    <h3>{{ subjectTitle(subject) }}</h3>
    <p>{{ subject.title }}</p>
  </base-card>
</template>
<script>
import BaseCard from "@/components/UI/BaseCard.vue";

export default {
  components: {BaseCard},
  props: ["subject"],
  methods: {
    subjectTitle(subject) {
      return subject.id + "@" + subject.title;
    }
  }};
</script>