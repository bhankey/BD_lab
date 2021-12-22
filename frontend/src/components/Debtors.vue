<template>
  <div class="debtors">
    <div class="field">Account ID</div>
    <div class="field">Sum of debts</div>
    <div class="field">Income for all time</div>
    <div class="field">Outgo for all time</div>
  </div>
  <div v-for="(item, idx) in debtors" :key="idx">
    <div class="debtors">
      <div class="field block with_separator">
        {{ item.account_id }}
      </div>
      <div class="field block with_separator">
        {{ item.debt_sum }}
      </div>
      <div class="field block with_separator">
        {{ item.income }}
      </div>
      <div class="field block with_separator">
        {{ item.outgo }}
      </div>
    </div>
    <div class="line"></div>
  </div>
</template>


<script>
export default {
  name: "Debtors",
  props: {
    debtors: Array,
  },
};
</script>

<style scoped>
.debtors {
  display: grid;
  grid-auto-rows: max-content;
  grid-template-columns: 1fr 1fr 1fr 1fr;
  margin-bottom: 20px;
  gap: 15px;
}

/* .block {
  background: rgba(255, 255, 255, 0.1);
  padding: 20px 0;
  text-align: center;
  position: relative;
}
.line {
  border-bottom: 3px solid rgb(0, 0, 0); /* Параметры линии 
}

.block.with_separator:after {
  content: "";
  position: absolute;
  top: 0;
  right: -12px;
  height: 100%;
  width: 3px;
  background: black;
}
 */
</style>