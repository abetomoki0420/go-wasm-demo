function App() {
  const onClick = async () => {
    console.log(
      window.add(20, 30)
      )
    console.log(
      window.sum([1,2,3,4,5])
      )
  }

  return (
    <div>
      <h1>Go Wasm</h1>
      <button onClick={onClick}>Run</button>
    </div>
  )
}

export default App
