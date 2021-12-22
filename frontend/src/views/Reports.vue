<template>
  <div class="home">
    <div v-if="isError" class="error">
      {{ errorMessage }}
    </div>
    <div class="button_wrapper">
      <div class="enter_wrapper">
        <label for="year">Year</label>
        <input v-model="year" @input="isError = false" type="number" required />
        <button @click="showTurnOverSheets()" class="save button">
          Show turn-over report
        </button>
      </div>
      <div class="enter_wrapper">
        <div class=""></div>
        <div class=""></div>
        <button @click="showDebtors()" class="save button">Get debtors</button>
      </div>
    </div>

    <div v-if="turnOverSheets != null" class="">
      <div class="table_row">
        <div>Account ID</div>
        <div>Starting</div>
        <div>Final</div>
        <div>January</div>
        <div>February</div>
        <div>March</div>
        <div>April</div>
        <div>May</div>
        <div>June</div>
        <div>July</div>
        <div>August</div>
        <div>September</div>
        <div>October</div>
        <div>November</div>
        <div>December</div>
      </div>

      <div class="table">
        <div class="table_row" v-for="item in turnOverSheets" :key="item.id">
          <div>
            {{ item.account_id }}
          </div>
          <div>
            {{ item.start_sum?.toFixed(2) }}
          </div>
          <div>
            {{ item.end_sum?.toFixed(2) }}
          </div>
          <div class="" v-for="(month, idx) in item.month_report" :key="idx">
            <TurnOverMonth
              v-bind:income="month.income ?? 0"
              v-bind:sum="month.sum ?? 0"
              v-bind:outgo="month.outgo ?? 0"
            >
            </TurnOverMonth>
          </div>
        </div>
      </div>
    </div>
    <Debtors v-if="debtors != null" v-bind:debtors="debtors"> </Debtors>
  </div>
</template>

<script>
// create| change |delete | show all

// @ is an alias to /src
import TurnOverMonth from "@/components/TurnoverMonth.vue";
import Debtors from "@/components/Debtors.vue";

const Month = [
  "January",
  "February",
  "March",
  "April",
  "May",
  "June",
  "July",
  "August",
  "September",
  "October",
  "November",
  "December",
];

export default {
  name: "Payments",
  components: {
    TurnOverMonth,
    Debtors,
  },

  data() {
    return {
      account_ids: null,
      year: null,
      reason: null,
      sum: null,
      editIndex: null,
      turnOverSheets: null,
      debtors: null,
      isError: false,
      errorMessage: "",
    };
  },

  methods: {
    showTurnOverSheets() {
      console.log(this.new_name);
      fetch("http://localhost:9005/reports/turnover_sheets", {
        method: "POST",

        body: JSON.stringify({
          year: this.year,
          account_ids: this.account_ids,
        }),

        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
      })
        .then((res) => res.json())
        .then((data) => {
          console.log(data);
          if (data.error) {
            this.isError = true;
            this.errorMessage = data.error;
          } else {
            this.debtors = null;
            this.turnOverSheets = data;
          }
        })
        .catch((error) => console.log("error", error));
    },

    showDebtors() {
      console.log(this.new_name);
      fetch("http://localhost:9005/reports/debtors", {
        method: "POST",

        body: JSON.stringify({
          account_ids: this.account_ids,
        }),

        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
      })
        .then((res) => res.json())
        .then((data) => {
          console.log(data);
          if (data.error) {
            this.isError = true;
            this.errorMessage = data.error;
          } else {
            this.debtors = data;
            this.turnOverSheets = null;
          }
        })
        .catch((error) => console.log("error", error));
    },
  },
};
</script>

<style scoped>
.table {
  text-align: center;
}

.block {
  background: rgba(255, 255, 255, 0.1);
  padding: 20px 0;
  text-align: center;
  position: relative;
}
.line {
  border-bottom: 3px solid rgb(0, 0, 0); /* Параметры линии */
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

.enter_wrapper {
  display: grid;
  gap: 10px;
  grid-template-rows: 1fr 1fr 1fr;
  margin-bottom: 10px;
  width: 200px;
  margin-right: 10px;
}

.table_row {
  display: grid;
  grid-auto-rows: max-content;
  grid-template-columns: 0.5fr 0.5fr 0.5fr 0.5fr 0.5fr 0.5fr 0.5fr 0.5fr 0.5fr 0.5fr 0.5fr 0.5fr 0.5fr 0.5fr 0.5fr;
  margin-bottom: 20px;
  /* gap: 15px; */
}

.error {
  background-color: coral;
  color: white;
  font-weight: bold;
}

.model {
  position: absolute;
  top: 0;
  bottom: 0;
  right: 0;
  left: 0;
  background-color: rgba(230, 230, 230, 0.8);
  display: grid;
  justify-content: center;
  align-items: center;
}

.create {
  margin-bottom: 30px;
}

.one_enter {
  justify-self: right;
  justify-content: right;
  text-align: right;
  align-items: right;
  grid-column: -1;
}

.button {
  background-color: #1e201e; /* Green */
  border: none;
  color: white;
  padding: 5px 5px;
  text-align: center;
  text-decoration: none;
  display: inline-block;
  font-size: 16px;
}

.home {
  width: 90%;
  margin: auto;
}

.button_wrapper {
  display: grid;
  grid-template-columns: 1.5fr 1.5fr;
  width: min-content;
}
</style>