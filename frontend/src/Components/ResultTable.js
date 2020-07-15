import React, { useState } from "react";
import { Table, Container } from "semantic-ui-react";

const defaultValue = () => [
  {
    orderName: "ABC",
    custCompany: "XYZ",
    custName: "def",
    orderDate: "1/2/3",
    deliveredAmount: 12,
    totalAmount: 23,
  },
  {
    orderName: "ABCD",
    custCompany: "XYDZ",
    custName: "defS",
    orderDate: "1/22/3",
    deliveredAmount: 112,
    totalAmount: 232,
  },
];
// const ResultTable = (props) => {
//   props.map((order) => [
//     {
//       orderName: order.orderName,
//       custCompany: order.custCompany,
//       custName: order.custName,
//       orderDate: order.orderDate,
//       deliveredAmount: order.deliveredAmount,
//       totalAmount: order.totalAmount,
//     },
//   ]);
// };
const OrderTable = () => {
  const [orders, setOrders] = useState(defaultValue);
  return (
    <div style={(styles.bottomMargin, styles.textCenter, styles.padding)}>
      <Table celled>
        <Table.Header>
          <Table.Row>
            <Table.HeaderCell>Order Name</Table.HeaderCell>
            <Table.HeaderCell>Customer Company</Table.HeaderCell>
            <Table.HeaderCell>Customer Name</Table.HeaderCell>
            <Table.HeaderCell>Order Date</Table.HeaderCell>
            <Table.HeaderCell>Delivered Amount</Table.HeaderCell>
            <Table.HeaderCell>Total Amount</Table.HeaderCell>
          </Table.Row>
        </Table.Header>

        <Table.Body>
          {orders.map((order, index) => (
            <Table.Row key={index}>
              <Table.Cell>{order.orderName}</Table.Cell>
              <Table.Cell>{order.custCompany}</Table.Cell>
              <Table.Cell>{order.custName}</Table.Cell>
              <Table.Cell>{order.orderDate}</Table.Cell>
              <Table.Cell>{order.deliveredAmount}</Table.Cell>
              <Table.Cell>{order.totalAmount}</Table.Cell>
            </Table.Row>
          ))}
        </Table.Body>
      </Table>
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
  padding: {
    leftPadding: "1em",
    rightPadding: "1em",
  },
};

export default OrderTable;
