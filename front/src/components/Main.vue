<template>
  <div class="mainComponent">
    <div class="search-folder" style="text-align: center;">
      <el-input size="large" placeholder="検索するワードを入力してください" v-model="searchWord" class="input-with-select" style="width: 300px; text-align: center">
        <el-select slot="prepend" label="hitlimit" placeholder="表示する検索ヒット数">
          <el-option label="表示する検索ヒット数"></el-option>
          <el-input-number v-model="limit" controls-position="left" :min="1" :max="30"></el-input-number>
        </el-select>
        <el-button type="primary" :loading="isLoading" v-on:click="search" slot="append" icon="el-icon-search"></el-button>
      </el-input>
    </div>
    <el-table
      :data="searchedData"
      style="width: 1000px; text-align: center;"
      empty-text="no data">
      <el-table-column
        label="title"
        width="300">
        <template slot-scope="scope">
          <a :href="scope.row.url"> {{ scope.row.title }} </a>
        </template>
      </el-table-column>
      <el-table-column
        prop="desc"
        label="description">
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
      searchedData: [{ id: 0, url: "", desc: "", title: "" }],
      limit: 10,
      isLoading: false
    };
  },
  methods: {
    search: function() {
      this.isLoading = true;
      let url = "http://localhost:8080/api/v1/page";
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

.search-folder {
}
.mainComponent {
  width: 100%;
  height: 100%;
  text-align: center;
  display: flex;
  flex-direction: column;
  align-items: center;
}
</style>
