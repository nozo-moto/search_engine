<template>
  <div class="mainComponent">
    <el-input placeholder="input search word" v-model="searchWord"></el-input>
    <el-input placeholder="limit" v-model="limit"></el-input>
    <el-button type="primary" :loading="isLoading" v-on:click="search">Search</el-button>
    <el-table
      :data="searchedData"
      style="width: 100%">
      <el-table-column
        prop="url"
        label="url"
        width="180">
      </el-table-column>
      <el-table-column
        prop="desc"
        label="description"
        width="180">
      </el-table-column>
      <el-table-column
        prop="title"
        label="title">
      </el-table-column>
    </el-table>
  </div>
</template>

<script>
export default {
  name: "mainComponent",
  data() {
    return {
      searchWord: "",
      searchedData: [
        { id: 0, url: "hoge.com", desc: "test desc", title: "test title" }
      ],
      limit: 0,
      isLoading: false
    };
  },
  methods: {
    search: function() {
      this.isLoading = true;
      let url = "localhost:8080/api/v1/page";
      url += "?q=" + this.searchWord + "&limit=" + this.limit;
      fetch(url, {
        method: "GET"
      })
        .then(response => {
          if (response.ok) {
            return response.json();
          }
          throw new Error(response);
        })
        .then(res_json => {
          // TODO
          // Table
          // eslint-disable-next-line
          console.log(res_json);
          this.isLoading = false;
          this.searchedData = res_json;
        })
        .catch(err => {
          // eslint-disable-next-line
          console.error(err);
        });
    }
  }
};
</script>

<!-- Add "scoped" attribute to limit CSS to this component only -->
<style scoped>
h3 {
  margin: 40px 0 0;
}
ul {
  list-style-type: none;
  padding: 0;
}
li {
  display: inline-block;
  margin: 0 10px;
}
a {
  color: #42b983;
}
</style>
