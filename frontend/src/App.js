import React from "react";
import { Container } from "semantic-ui-react";
import Search from "./Components/Search";
import Table from "./Components/ResultTable";
import "./App.css";

const App = () => (
  <Container style={styles.containerPadding}>
    <Search />
    <Table />
  </Container>
);

const styles = {
  containerPadding: {
    paddingTop: "2em",
    paddingBottom: "2em",
  },
};
export default App;
