import React from "react";
import { Grid, Button } from "@material-ui/core";

export default function Toolbar() {
  return (
    <Grid container justifyContent="center">
      <Grid item xs={10}>
        <Grid container>
          <Button variant="contained" color="secondary">
            EXECUTE
          </Button>
        </Grid>
      </Grid>
    </Grid>
  );
}
