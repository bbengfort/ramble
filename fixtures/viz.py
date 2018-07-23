#!/usr/bin/env python

import argparse

import numpy as np
import pandas as pd
import seaborn as sns
import matplotlib.pyplot as plt

sns.set_style('whitegrid')
sns.set_context('notebook')


def draw_benchmark(path):
    df = pd.read_csv(path)
    _, ax = plt.subplots(ncols=2, figsize=(18,6), sharey=True)

    means = df.groupby('clients')['throughput'].mean()
    std = df.groupby('clients')['throughput'].std()

    ax[0].plot(means, label="throughput")
    ax[0].fill_between(np.arange(1, 13), means+std, means-std, alpha=0.25)

    ax[0].set_xlim(1,12)
    ax[0].set_ylabel("messages/second")
    ax[0].set_xlabel("concurrent clients")
    ax[0].set_title("Ramble Benchmark")
    ax[0].legend(frameon=True)

    sns.barplot('clients', 'throughput', ax=ax[1], data=df, palette="GnBu_d")
    ax[1].set_ylabel("")
    ax[1].set_xlabel("concurrent clients")
    ax[1].set_title("Ramble Benchmark")

    plt.tight_layout()
    plt.savefig("benchmark.png")

if __name__ == '__main__':
    parser = argparse.ArgumentParser(
        description="draw the benchmark visualization from a dataset",
    )

    parser.add_argument("data")

    args = parser.parse_args()
    draw_benchmark(args.data)
