import React, { ChangeEvent, useState } from "react";
import { Button, TextField, Grid, Box } from "material-ui-core";

export default function Playground() {
  const [query, setQuery] = useState(`CREATE TABLE todos (
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    content VARCHAR(255) NOT NULL
);
INSERT INTO todos ( content )
VALUES
( 'todo1' ),
( 'todo2' ),
( 'todo3' ),
( 'todo4' );
SELECT *
FROM todos;`);

  const handleQueryChange = (e: ChangeEvent<HTMLInputElement>) => {
    setQuery(e.target.value);
  };

  const [result, setResult] = useState("");
  const [executeButtonEnabled, setExecuteButtonEnabled] = useState(true);

  const handleExecuteButtonClick = async () => {
    setExecuteButtonEnabled(false);
    setResult("Executing...");

    const sleep = (ms: number) =>
      new Promise((resolve) => setTimeout(resolve, ms));
    await sleep(1000);

    setExecuteButtonEnabled(true);
    setResult(`Finished!
----------------------------------------------------------------------------------------------------
+--------------------------------------+---------+
| id                                   | content |
+--------------------------------------+---------+
| cc944779-52af-435b-beb9-b8539cdf2176 | todo1   |
| 3bde42c1-a40c-4445-8f80-092a998c115f | todo2   |
| 8d37a77f-ba18-4a27-b1b0-9ed6b8824d96 | todo3   |
| 37e492a8-da0e-413d-bf11-5c344ef04f37 | todo4   |
+--------------------------------------+---------+
`);
  };

  return (
    <Box className="playground" style={{ margin: 8 }}>
      <Grid container justifyContent="center" spacing={1}>
        <Grid item xs={8}>
          <Grid container>
            <Button
              variant="contained"
              color="secondary"
              disabled={!executeButtonEnabled}
              onClick={handleExecuteButtonClick}
            >
              EXECUTE
            </Button>
          </Grid>
        </Grid>

        <Grid item xs={8}>
          <div>
            <TextField
              variant="outlined"
              multiline
              rows={20}
              fullWidth={true}
              value={query}
              onChange={handleQueryChange}
            />
          </div>
        </Grid>

        <Grid item xs={8}>
          <TextField
            variant="outlined"
            multiline
            rows={20}
            fullWidth={true}
            InputProps={{ readOnly: true }}
            value={result}
          />
        </Grid>
      </Grid>
    </Box>
  );
}
