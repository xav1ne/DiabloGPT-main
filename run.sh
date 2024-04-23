python3 maqp.py --generate_hdf --generate_sampled_hdfs --generate_ensemble
python3 maqp.py --evaluate_cardinalities --rdc_spn_selection --max_variants 1 --pairwise_rdc_path ../imdb-benchmark/spn_ensembles/pairwise_rdc.pkl --dataset imdb-light --target_path ./baselines/cardinality_estimation/results/deepdblight/imdb_light_model_based_budget_5.csv --ensemble_location ../imdb-benchmark/spn_ensembles/ensemble_join_3_budget_5_10000000.pkl --query_file_location ./benchmarks/job-light/sql/test-num.sql --ground_truth_file_location ./benchmarks/job-light/sql/true_cardinalities.csv

# Path: run.sh