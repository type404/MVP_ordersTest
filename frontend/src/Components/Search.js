import React, { useState, useEffect } from "react";
import { Header, Button, Divider, Input } from "semantic-ui-react";
import ResultTable from "./ResultTable";

const dummyData = [
  {
    orderName: "qwe",
    custCompany: "XYZ",
    custName: "def",
    orderDate: "1/2/3",
    deliveredAmount: 12,
    totalAmount: 23,
  },
  {
    orderName: "qaz",
    custCompany: "XYDZ",
    custName: "defS",
    orderDate: "1/22/3",
    deliveredAmount: 112,
    totalAmount: 232,
  },
];

const OrderSearch = () => {
  const [searchTerm, setSearchTerm] = useState("");
  const [searchResults, setSearchResults] = useState(dummyData);
  const handleChange = (e) => {
    setSearchTerm(e.target.value.toLowerCase());
  };
  const updateList = () => {
    const results = dummyData.filter((data) =>
      data.orderName.toLowerCase().includes(searchTerm)
    );
    setSearchResults(results);
  };
  return (
    <div>
      <Header as="h1" style={styles.textCenter}>
        Search Order
      </Header>
      <Divider />
      <div style={styles.bottomMargin}>
        <Input
          fluid
          style={(styles.bottomMargin, styles.textCenter)}
          value={searchTerm}
          onChange={handleChange}
        />
      </div>
      <Divider />
      <div style={(styles.bottomMargin, styles.textCenter)}>
        <Button primary onClick={updateList}>
          Search
        </Button>
      </div>
    </div>
  );
};
const styles = {
  bottomMargin: {
    marginBottom: "1em",
  },
  textCenter: {
    textAlign: "center",
  },
};
export default OrderSearch;
