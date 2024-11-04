import tkinter as tk
from tkinter import filedialog
from tkinter.ttk import Scale, Button
import time
import threading
import matplotlib.pyplot as plt
from matplotlib.backends.backend_tkagg import FigureCanvasTkAgg
from mpl_toolkits.mplot3d import Axes3D
import numpy as np

class CubeReplayApp:
    def __init__(self, root):
        self.root = root
        self.root.title("5x5x5 Magic Cube Replay")
        self.is_playing = False
        self.speed = 5.0
        self.current_step = 0
        self.steps = []
        self.cube_size = 5  # Default cube size

        # Load State Button
        self.load_button = Button(root, text="Load Experiment File", command=self.load_file)
        self.load_button.pack()

        # Play/Pause Button
        self.play_button = Button(root, text="Play", command=self.toggle_play)
        self.play_button.pack()

        # Progress Bar
        self.progress = tk.Scale(root, from_=0, to=100, orient="horizontal", command=self.seek)
        self.progress.pack(fill="x", padx=10)

        # Speed Control
        self.speed_control = Scale(root, from_=0.5, to=10.0, value=1.0, orient="horizontal", command=self.set_speed)
        self.speed_control.pack(fill="x", padx=10)

        # Label to display the current step's N value
        self.n_value_label = tk.Label(root, text="N = ")
        self.n_value_label.pack()

        # 3D Visualization using matplotlib
        self.figure = plt.Figure(figsize=(5, 5))
        self.ax = self.figure.add_subplot(111, projection="3d")
        self.canvas = FigureCanvasTkAgg(self.figure, master=root)
        self.canvas.get_tk_widget().pack()

        # Dictionary to hold scatter plot references for each point
        self.plot_points = {}

    def load_file(self):
        file_path = filedialog.askopenfilename()
        if file_path:
            with open(file_path, 'r') as file:
                lines = file.readlines()
                lines = [x.strip() for x in lines if x.strip()]

                # Initialize a 3D array for a 5x5x5 cube to store values from the file
                self.cube_state = np.zeros((5, 5, 5), dtype=int)

                # Parse the file into the 3D cube array
                for i in range(5):
                    for j in range(5):
                        one_line = lines[5 * i + j].split(' ')
                        for k in range(5):
                            self.cube_state[i][j][k] = int(one_line[k])

                # Parse steps
                self.steps = [line.strip() for line in lines[25:]]
                self.progress.configure(to=len(self.steps) - 1)

                # Initial rendering of the entire cube
                self.display_cube()

    def display_cube(self):
        self.ax.clear()
        self.plot_points.clear()  # Clear previous points

        # Plot each point with a neutral color and add text annotation for each point with its value
        for x in range(self.cube_size):
            for y in range(self.cube_size):
                for z in range(self.cube_size):
                    scatter_point = self.ax.scatter(x, y, z, color="none", s=100, edgecolor="gray")
                    self.plot_points[(x, y, z)] = scatter_point
                    self.ax.text(x, y, z, f"{self.cube_state[x][y][z]}", color="black", fontsize=8, ha='center')

        # Set plot limits and labels
        self.ax.set_xlim(0, self.cube_size - 1)
        self.ax.set_ylim(0, self.cube_size - 1)
        self.ax.set_zlim(0, self.cube_size - 1)
        self.ax.set_xlabel("X")
        self.ax.set_ylabel("Y")
        self.ax.set_zlabel("Z")
        self.canvas.draw()

    def toggle_play(self):
        if not self.is_playing:
            self.is_playing = True
            self.play_button.config(text="Pause")
            self.play_thread = threading.Thread(target=self.play)
            self.play_thread.start()
        else:
            self.is_playing = False
            self.play_button.config(text="Play")

    def play(self):
        while self.is_playing and self.current_step < len(self.steps) - 1:
            self.display_step(self.current_step)
            self.current_step += 1
            self.progress.set(self.current_step)
            time.sleep(1 / self.speed)
        self.is_playing = False
        self.play_button.config(text="Play")

    def seek(self, position):
        self.current_step = int(float(position))
        self.display_step(self.current_step)

    def set_speed(self, value):
        self.speed = float(value)

    def display_step(self, step_index):
        if step_index < len(self.steps):
            step = self.steps[step_index]
            # Parse each step assuming format "x1 y1 z1 x2 y2 z2 n"
            parts = step.split()
            if len(parts) == 7:
                x1, y1, z1, x2, y2, z2, n = map(int, parts)

                # Update the N value label
                self.n_value_label.config(text=f"N = {n}")

                # Highlight start and end points in different colors
                self.highlight_swap(x1, y1, z1, x2, y2, z2)

                # Update the cube state by swapping values
                self.cube_state[x1][y1][z1], self.cube_state[x2][y2][z2] = (
                    self.cube_state[x2][y2][z2],
                    self.cube_state[x1][y1][z1],
                )

                # Update the annotations with new values
                self.update_annotations()
                self.canvas.draw()

    def highlight_swap(self, x1, y1, z1, x2, y2, z2):
        # Reset all points to neutral color first
        for scatter in self.plot_points.values():
            scatter.set_color("gray")

        # Highlight the swapped points
        self.plot_points[(x1, y1, z1)].set_color("red")
        self.plot_points[(x2, y2, z2)].set_color("blue")

    def update_annotations(self):
        # Remove existing text annotations
        for text in self.ax.texts:
            text.remove()

        # Add updated text annotations with new values after the swap
        for x in range(self.cube_size):
            for y in range(self.cube_size):
                for z in range(self.cube_size):
                    # Add updated text annotations
                    self.ax.text(x, y, z, f"{self.cube_state[x][y][z]}", color="black", fontsize=8, ha='center')

root = tk.Tk()
app = CubeReplayApp(root)
root.mainloop()
