<template>
  <div class="home">
    <div v-if="isError" class="error">
      {{ errorMessage }}
    </div>
    <div class="create">
      <label for="new_name"> Enter name </label>
      <input v-model="new_name" @input="isError = false" type="text" required />
      <button @click="saveAccount()" class="save button">Create</button>
    </div>

    <div class="table">
      <div class="table_row">
        <div>ID</div>
        <div>Name</div>
        <div>Sum</div>
        <div class=""></div>
        <div class=""></div>
      </div>
      <div class="table_row" v-for="(item, idx) in accounts" :key="item.id">
        <div>
          {{ item.id }}
        </div>
        <div>
          {{ item.name }}
        </div>
        <div>
          {{ item.sum.toFixed(4) }}
        </div>
        <div class="">
          <button @click="editAccount(idx)" class="button">Edit</button>
        </div>
        <div class="">
          <button @click="deleteAccount(idx)" class="button">Delete</button>
        </div>
      </div>
    </div>
    <div v-if="isModelShow" class="model">
      <div class="edit">
        <div class="">Name</div>
        <input v-model="updateName" type="text" />
        <button @click="updateAccount" class="button">Save</button>
      </div>
    </div>
  </div>
</template>

<script>
// create| change |delete | show all

// @ is an alias to /src

export default {
  el: "#score",
  name: "Accounts",
  components: {},

  data() {
    return {
      editIndex: null,
      accounts: [],
      new_name: "",
      isError: false,
      errorMessage: "",
      updateName: "",
      isModelShow: false,
    };
  },

  methods: {
    saveAccount() {
      console.log(this.new_name);
      fetch("http://localhost:9005/account/create", {
        method: "POST",
        body: JSON.stringify({
          name: this.new_name,
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
            this.new_name = null;
            fetch("http://localhost:9005/account/get_all", {
              method: "GET",
            })
              .then((res) => res.json())
              .then((data) => {
                this.accounts = data;
              })
              .catch((error) => {
                console.log(error);
              });
          }
        })
        .catch((error) => console.log("error", error));
    },

    editAccount(id) {
      this.editIndex = id;
      this.isModelShow = true;

      this.updateName = this.accounts[id].name;
    },

    updateAccount() {
      console.log(this.new_name);
      fetch("http://localhost:9005/account/update", {
        method: "POST",

        body: JSON.stringify({
          account_id: this.accounts[this.editIndex].id,
          name: this.updateName,
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
            this.accounts[this.editIndex].name = this.updateName;
          }
          this.isModelShow = false;
        })
        .catch((error) => console.log("error", error));
    },

    deleteAccount(id) {
      if (confirm("Are you sure?")) {
        console.log("lol");
        fetch("http://localhost:9005/account/delete", {
          method: "DELETE",
          body: JSON.stringify({
            account_id: this.accounts[id].id,
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
              this.accounts.splice(id, 1);
            }
          })
          .catch((error) => console.log("error", error));
      }
    },
  },

  beforeMount() {
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

.table_row {
  display: grid;
  grid-auto-rows: max-content;
  grid-template-columns: 1fr 3fr 1fr 0.3fr 0.3fr;
  margin-bottom: 20px;
  gap: 15px;
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
  display: grid;
  gap: 10px;
  margin-bottom: 30px;
  margin-top: 30px;
  grid-template-columns: 1fr 1fr 1fr;
}

.hide {
  visibility: hidden !important;
}

.home {
  width: 80%;
  margin: auto;
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
</style>