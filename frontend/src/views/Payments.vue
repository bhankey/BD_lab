<template>
  <div class="home">
    <div v-if="isError" class="error">
      {{ errorMessage }}
    </div>
    <div class="enter_wrapper">
      <label for="">Account ID</label>
      <label for="reason">Reason</label>
      <label for="sum">Sum</label>

      <select v-model="accountID" @change="clientPayments">
        <option
          v-for="account in accounts"
          v-bind:value="account.id"
          v-bind:key="account.id"
        >
          {{ account.id }}
        </option>
      </select>
      <input v-model="reason" @input="isError = false" type="text" required />
      <input v-model="sum" @input="isError = false" type="number" required />
    </div>
    <button @click="savePayments()" class="save button">Create</button>

    <div class="table">
      <div class="table_row">
        <div>Account ID</div>
        <div>Payment reason</div>
        <div>Sum</div>
      </div>

      <div class="table_row" v-for="item in payments" :key="item.id">
        <div>
          {{ item.account_id }}
        </div>
        <div>
          {{ item.reason }}
        </div>
        <div>
          {{ item.sum?.toFixed(4) }}
        </div>
      </div>
    </div>
  </div>
</template>

<script>
// create| change |delete | show all

// @ is an alias to /src

export default {
  name: "Payments",
  components: {},

  data() {
    return {
      accounts: null,
      accountID: null,
      reason: null,
      sum: null,
      editIndex: null,
      payments: [],
      isError: false,
      errorMessage: "",
      isModelShow: false,
    };
  },

  methods: {
    savePayments() {
      console.log(this.new_name);
      fetch("http://localhost:9005/payments/create", {
        method: "POST",

        body: JSON.stringify({
          account_id: this.accountID,
          reason: this.reason,
          sum: this.sum,
        }),

        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
      })
        .then((res) => res.json())
        .then((data) => {
          if (!data.success) {
            this.isError = true;
            this.errorMessage = data.error;
          } else {
            this.accountID = null;
            this.reason = null;
            this.sum = null;

            fetch("http://localhost:9005/payments/get_all", {
              method: "GET",
            })
              .then((res) => res.json())
              .then((data) => {
                console.log(data);
                this.payments = data;
              })
              .catch((error) => {
                console.log(error);
              });
          }
        })
        .catch((error) => console.log("error", error));
    },
    clientPayments() {
      fetch("http://localhost:9005/payments/get_client_payments", {
        method: "POST",

        body: JSON.stringify({
          account_id: this.accountID,
        }),

        headers: {
          Accept: "application/json",
          "Content-Type": "application/json",
        },
      })
        .then((res) => res.json())
        .then((data) => {
          console.log(data);
          this.payments = data;
        })
        .catch((error) => {
          console.log(error);
        });
      fetch("http://localhost:9005/account/get_all", {
        method: "GET",
      })
        .then((res) => res.json())
        .then((data) => {
          console.log(data);
          this.accounts = data;
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },

  beforeMount() {
    fetch("http://localhost:9005/payments/get_all", {
      method: "GET",
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        this.payments = data;
      })
      .catch((error) => {
        console.log(error);
      });
    fetch("http://localhost:9005/account/get_all", {
      method: "GET",
    })
      .then((res) => res.json())
      .then((data) => {
        console.log(data);
        this.accounts = data;
      })
      .catch((error) => {
        console.log(error);
      });
  },
};
</script>

<style scoped>
.table {
  text-align: left;
}

.enter_wrapper {
  display: grid;
  grid-template-columns: 1fr 1fr 1fr;

  gap: 10px 10px;
}

.table_row {
  display: grid;
  grid-auto-rows: max-content;
  grid-template-columns: 1fr 3fr 1fr;
  margin-bottom: 20px;
  gap: 15px;
}

.error {
  background-color: coral;
  color: white;
  font-weight: bold;
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
  width: 100%;
  margin-bottom: 10px;
  margin-top: 10px;
}

.home {
  width: 80%;
  margin: auto;
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
</style>