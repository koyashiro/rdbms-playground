import React from "react";
import { TextField } from "@material-ui/core";

export default function Query() {
  return (
    <div>
      <TextField variant="outlined" multiline rows={30} fullWidth={true} />
    </div>
  );
}
