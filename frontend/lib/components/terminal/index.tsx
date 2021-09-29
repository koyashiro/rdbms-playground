import {
  ChangeEventHandler,
  createRef,
  FC,
  KeyboardEventHandler,
  useEffect,
  useState,
} from "react";

type Props = {
  command: (cmd: string) => Promise<string>;
};

const defaultHistory = `
Welcome to Postgres Playground!
`;

const commandLinePrefix = "> ";

const Terminal: FC<Props> = (props) => {
  const [mergedHistory, setMergedHistory] = useState<string>(defaultHistory);
  const [commandHistory, setCommandHistory] = useState<string[]>([]);
  const [textareaContent, setTextareaContent] = useState<string>("");
  const [textareaRows, setTextareaRows] = useState<number>(1);
  const [commandHistoryInvertedIndex, setCommandHistoryInvertedIndex] =
    useState<number>(0);

  const textareaRef = createRef<HTMLTextAreaElement>();

  useEffect(
    () => textareaRef.current?.scrollIntoView(),
    [textareaRef, textareaRows]
  );

  const addHistory = (command: string, res: string) => {
    const fixedCommand = command
      .split("\n")
      .join(`\n${" ".repeat(commandLinePrefix.length)}`);
    const line = `${commandLinePrefix}${fixedCommand}\n${res}\n`;

    setCommandHistory(commandHistory.concat(command));
    setMergedHistory(mergedHistory + line);
  };

  const loadHistory = (newIndex: number) => {
    if (newIndex >= 1 && newIndex <= commandHistory.length) {
      setCommandHistoryInvertedIndex(newIndex);
      setTextareaContent(commandHistory[commandHistory.length - newIndex]);
    } else if (newIndex <= 0) {
      setCommandHistoryInvertedIndex(newIndex);
      setTextareaContent("");
    }
  };

  const handleChange: ChangeEventHandler<HTMLTextAreaElement> = (e) => {
    setTextareaContent(e.currentTarget.value);
  };

  const handleKeyPress: KeyboardEventHandler<HTMLTextAreaElement> = async (
    e
  ) => {
    switch (e.key) {
      case "Enter":
        if (textareaContent.trimEnd().endsWith(";")) {
          e.preventDefault();

          const res = await props.command(textareaContent);
          addHistory(textareaContent, res);

          setTextareaContent("");
          setTextareaRows(1);
        } else {
          setTextareaRows(textareaRows + 1);
        }
        break;
      case "ArrowUp":
        loadHistory(commandHistoryInvertedIndex + 1);
        break;
      case "ArrowDown":
        loadHistory(commandHistoryInvertedIndex - 1);
        break;
      default:
        break;
    }
  };

  return (
    <>
      <div className="flex flex-col w-full h-full overflow-x-hidden overflow-y-scroll font-mono text-base text-green-500 bg-black">
        <pre>{mergedHistory}</pre>
        <div className="flex w-full">
          <pre className="flex-initial">{commandLinePrefix}</pre>
          <textarea
            className="flex-auto outline-none bg-black text-green-500 font-mono text-base"
            rows={textareaRows}
            onChange={handleChange}
            onKeyDown={handleKeyPress}
            value={textareaContent}
            ref={textareaRef}
          />
        </div>
      </div>
    </>
  );
};

export default Terminal;
