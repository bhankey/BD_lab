
<template>
  <div class="month">
    <div class="field">
      {{ income?.toFixed(2) }}
    </div>
    <div class="field">
      {{ outgo?.toFixed(2) }}
    </div>
    <div class="field">
      {{ sum?.toFixed(2) }}
    </div>
  </div>
</template>


<script>
export default {
  name: "TurnOverMonth",
  props: {
    income: Number,
    outgo: Number,
    sum: Number,
  },
};
</script>

<style scoped>
.month {
  display: grid;
  grid-template-rows: min-content min-content min-content;
}
</style>