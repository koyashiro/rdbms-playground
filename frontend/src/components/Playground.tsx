import React from "react";
import { Grid } from "@material-ui/core";
import Query from "./Query";
import Result from "./Result";

export default function Playground() {
  return (
    <Grid container justifyContent="center" spacing={1}>
      <Grid item xs={5}>
        <Query />
      </Grid>
      <Grid item xs={5}>
        <Result />
      </Grid>
    </Grid>
  );
}
